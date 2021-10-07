package job

import v1client "github.com/hashicorp/nomad-openapi/clients/go/v1"

const (
	PackPathKey           = "pack-path"
	PackNameKey           = "pack-name"
	PackRegistryKey       = "pack-registry"
	PackDeploymentNameKey = "pack-deployment-name"
	PackJobKey            = "pack-job"
	PackVersionKey        = "pack-version"
)

// add metadata to the job for in cluster querying and management
func (r *Runner) setJobMeta(job *v1client.Job) {
	jobMeta := make(map[string]string)

	// If current job meta isn't nil, use that instead
	if job.Meta != nil {
		jobMeta = *job.Meta
	}

	// Add the Nomad Pack custom metadata.
	jobMeta[PackPathKey] = r.runnerCfg.PathPath
	jobMeta[PackNameKey] = r.runnerCfg.PackName
	jobMeta[PackRegistryKey] = r.runnerCfg.RegistryName
	jobMeta[PackDeploymentNameKey] = r.runnerCfg.DeploymentName
	jobMeta[PackJobKey] = *job.Name
	jobMeta[PackVersionKey] = r.runnerCfg.PackVersion

	// Replace the job metadata with our modified version.
	job.Meta = &jobMeta
}