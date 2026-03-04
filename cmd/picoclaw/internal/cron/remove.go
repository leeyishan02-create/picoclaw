// =============================================================================
// remove.go - 删除定时任务命令实现
// =============================================================================
// 这个文件实现了PicoClaw的删除定时任务功能。
// =============================================================================

package cron

// =============================================================================
// 导入 (Imports)
// =============================================================================
import "github.com/spf13/cobra"

// =============================================================================
// newRemoveCommand - 创建删除命令
// =============================================================================
// 创建"picoclaw cron remove"子命令（别名"rm"）。
// 根据任务ID删除定时任务。
//
// 参数：
//   - storePath: 返回定时任务存储文件路径的函数
//
// 位置参数：
//   - args[0]: 任务ID
//
// 使用示例：
//
//	picoclaw cron remove 1
func newRemoveCommand(storePath func() string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   "Remove a job by ID",
		Args:    cobra.ExactArgs(1), // 必须提供1个参数（任务ID）
		Example: `picoclaw cron remove 1`,
		RunE: func(_ *cobra.Command, args []string) error {
			cronRemoveCmd(storePath(), args[0])
			return nil
		},
	}

	return cmd
}
