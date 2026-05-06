package artifactregistry

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud artifactregistry command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "artifactregistry",
		Usage: "manage Artifact Registry API resources",
		Commands: []*cli.Command{
			{
				Name:  "apt-artifacts",
				Usage: "Manage apt-artifacts resources",
				Commands: []*cli.Command{
					{
						Name:  "import",
						Usage: "import apt-artifacts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing import...")
							return nil
						},
					},
				},
			},
			{
				Name:  "attachments",
				Usage: "Manage attachments resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "attachment", Usage: "The attachment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/attachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("attachment"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "attachment", Usage: "The attachment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/attachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("attachment"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "docker-images",
				Usage: "Manage docker-images resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list docker-images",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe docker-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "docker_image", Usage: "The docker_image.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/dockerImages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("docker_image"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "files",
				Usage: "Manage files resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "file", Usage: "The file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/files/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("file"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "file", Usage: "The file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/files/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("file"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "file", Usage: "The file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/files/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("file"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "maven-artifacts",
				Usage: "Manage maven-artifacts resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list maven-artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe maven-artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "maven_artifact", Usage: "The maven_artifact.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/mavenArtifacts/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("maven_artifact"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "npm-packages",
				Usage: "Manage npm-packages resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list npm-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe npm-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "npm_package", Usage: "The npm_package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/npmPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("npm_package"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "packages",
				Usage: "Manage packages resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "project-settings",
				Usage: "Manage project-settings resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe project-settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/projectSettings", cmd.String("project"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update project-settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/projectSettings", cmd.String("project"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "python-packages",
				Usage: "Manage python-packages resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list python-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe python-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "python_package", Usage: "The python_package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pythonPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("python_package"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "repositories",
				Usage: "Manage repositories resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "export-artifact",
						Usage: "export-artifact repositories",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export-artifact...")
							return nil
						},
					},
				},
			},
			{
				Name:  "rules",
				Usage: "Manage rules resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("rule"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("rule"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("rule"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "tags",
				Usage: "Manage tags resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list tags",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe tags",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"), cmd.String("tag"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete tags",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
				},
			},
			{
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
					{
						Name:  "batch-delete",
						Usage: "batch-delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"))
							fmt.Printf("Executing batch-delete on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The package.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("package"), cmd.String("version"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "vpcsc-config",
				Usage: "Manage vpcsc-config resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe vpcsc-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcscConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update vpcsc-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcscConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "yum-artifacts",
				Usage: "Manage yum-artifacts resources",
				Commands: []*cli.Command{
					{
						Name:  "import",
						Usage: "import yum-artifacts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing import...")
							return nil
						},
					},
				},
			},
		},
	}
}
