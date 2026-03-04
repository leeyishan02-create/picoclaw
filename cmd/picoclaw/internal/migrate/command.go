// =============================================================================
// migrate 包 - 数据迁移模块
// =============================================================================
// 这个包负责PicoClaw的数据迁移功能。
// 支持从其他AI助手（如OpenClaw）迁移配置和数据到PicoClaw。
// =============================================================================

package migrate

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/pkg/migrate" // 迁移核心功能
)

// =============================================================================
// NewMigrateCommand - 创建迁移命令
// =============================================================================
// 这是migrate模块的入口函数，创建数据迁移命令。
// 当用户运行"picoclaw migrate"时会调用此函数。
//
// 命令行参数：
//   - --from: 指定源系统（默认：openclaw）
//   - --dry-run: 试运行，不实际执行迁移
//   - --refresh: 刷新已迁移的数据
//   - --force: 强制覆盖现有数据
//
// 使用示例：
//
//	# 从OpenClaw迁移
//	picoclaw migrate
//
//	# 试运行
//	picoclaw migrate --dry-run
func NewMigrateCommand() *cobra.Command {
	var opts migrate.Options // 迁移选项

	// 创建迁移命令
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate from xxxclaw(openclaw, etc.) to picoclaw",
		Args:  cobra.NoArgs,
		Example: `  picoclaw migrate
  picoclaw migrate --from openclaw
  picoclaw migrate --dry-run
  picoclaw migrate --refresh
  picoclaw migrate --force`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			m := migrate.NewMigrateInstance(opts)
			result, err := m.Run(opts)
			if err != nil {
				return err
			}
			if !opts.DryRun {
				m.PrintSummary(result)
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&opts.DryRun, "dry-run", false,
		"Show what would be migrated without making changes")
	cmd.Flags().StringVar(&opts.Source, "from", "openclaw",
		"Source to migrate from (e.g., openclaw)")
	cmd.Flags().BoolVar(&opts.Refresh, "refresh", false,
		"Re-sync workspace files from OpenClaw (repeatable)")
	cmd.Flags().BoolVar(&opts.ConfigOnly, "config-only", false,
		"Only migrate config, skip workspace files")
	cmd.Flags().BoolVar(&opts.WorkspaceOnly, "workspace-only", false,
		"Only migrate workspace files, skip config")
	cmd.Flags().BoolVar(&opts.Force, "force", false,
		"Skip confirmation prompts")
	cmd.Flags().StringVar(&opts.SourceHome, "source-home", "",
		"Override source home directory (default: ~/.openclaw)")
	cmd.Flags().StringVar(&opts.TargetHome, "target-home", "",
		"Override target home directory (default: ~/.picoclaw)")

	return cmd
}
