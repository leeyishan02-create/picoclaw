// =============================================================================
// cron 包 - 定时任务管理模块
// =============================================================================
// 这个包负责PicoClaw的定时任务（cron jobs）管理功能。
// 用户可以创建、列出、删除、启用和禁用定时任务。
//
// 定时任务存储在配置工作区的 cron/jobs.json 文件中。
// =============================================================================

package cron

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"fmt"           // 格式化输出
	"path/filepath" // 路径操作

	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/cmd/picoclaw/internal" // 内部工具函数
)

// =============================================================================
// NewCronCommand - 创建定时任务命令
// =============================================================================
// 这是cron模块的入口函数，创建定时任务相关的子命令。
// 当用户运行"picoclaw cron"时会调用此函数。
//
// 子命令说明：
// - list (别名: l): 列出所有定时任务
// - add: 添加新的定时任务
// - remove (别名: rm): 删除定时任务
// - enable: 启用定时任务
// - disable: 禁用定时任务
//
// 工作原理：
// - 使用PersistentPreRunE在执行任何子命令前加载配置
// - 动态解析storePath，确保反映当前配置
// - 所有子命令共享同一个storePath
func NewCronCommand() *cobra.Command {
	var storePath string // 定时任务存储文件路径

	// 创建定时任务命令
	// Use: 命令名称
	// Aliases: 命令别名，"c"是"cron"的简写
	// Short: 简短描述
	// Args: 参数验证，cobra.NoArgs表示不接受任何位置参数
	// RunE: 默认行为，显示帮助信息
	// PersistentPreRunE: 在执行子命令前运行的函数，用于加载配置
	cmd := &cobra.Command{
		Use:     "cron",
		Aliases: []string{"c"},
		Short:   "Manage scheduled tasks",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
		// 在执行任何子命令前加载配置并解析storePath
		// 这样可以确保路径反映当前配置，并在所有子命令间共享
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := internal.LoadConfig()
			if err != nil {
				return fmt.Errorf("error loading config: %w", err)
			}
			// 定时任务存储在 工作区/cron/jobs.json
			storePath = filepath.Join(cfg.WorkspacePath(), "cron", "jobs.json")
			return nil
		},
	}

	// 添加子命令
	cmd.AddCommand(
		newListCommand(func() string { return storePath }),    // 列出命令
		newAddCommand(func() string { return storePath }),     // 添加命令
		newRemoveCommand(func() string { return storePath }),  // 删除命令
		newEnableCommand(func() string { return storePath }),  // 启用命令
		newDisableCommand(func() string { return storePath }), // 禁用命令
	)

	return cmd
}
