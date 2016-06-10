package db

import (
	"strconv"
	"strings"
)

func (db *SQLDB) GetAllStartedBuilds() ([]BuildDB, error) {
	rows, err := db.conn.Query(`
		SELECT ` + qualifiedBuildColumns + `
		FROM builds b
		LEFT OUTER JOIN jobs j ON b.job_id = j.id
		LEFT OUTER JOIN pipelines p ON j.pipeline_id = p.id
		LEFT OUTER JOIN teams t ON b.team_id = t.id
		WHERE b.status = 'started'
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bs := []BuildDB{}

	for rows.Next() {
		build, _, err := db.buildFactory.ScanBuild(rows)
		if err != nil {
			return nil, err
		}

		bs = append(bs, build)
	}

	return bs, nil
}

func (db *SQLDB) UpdateBuildPreparation(buildPrep BuildPreparation) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = db.buildPrepHelper.UpdateBuildPreparation(tx, buildPrep)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *SQLDB) ResetBuildPreparationsWithPipelinePaused(pipelineID int) error {
	_, err := db.conn.Exec(`
			UPDATE build_preparation
			SET paused_pipeline='blocking',
			    paused_job='unknown',
					max_running_builds='unknown',
					inputs='{}',
					inputs_satisfied='unknown'
			FROM build_preparation bp, builds b, jobs j
			WHERE bp.build_id = b.id AND b.job_id = j.id
				AND j.pipeline_id = $1 AND b.status = 'pending' AND b.scheduled = false
		`, pipelineID)
	return err
}

func (db *SQLDB) DeleteBuildEventsByBuildIDs(buildIDs []int) error {
	if len(buildIDs) == 0 {
		return nil
	}

	interfaceBuildIDs := make([]interface{}, len(buildIDs))
	for i, buildID := range buildIDs {
		interfaceBuildIDs[i] = buildID
	}

	indexStrings := make([]string, len(buildIDs))
	for i := range indexStrings {
		indexStrings[i] = "$" + strconv.Itoa(i+1)
	}

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(`
   DELETE FROM build_events
	 WHERE build_id IN (`+strings.Join(indexStrings, ",")+`)
	 `, interfaceBuildIDs...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE builds
		SET reap_time = now()
		WHERE id IN (`+strings.Join(indexStrings, ",")+`)
	`, interfaceBuildIDs...)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
