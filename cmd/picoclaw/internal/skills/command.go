// =============================================================================
// skills 包 - 技能管理模块
// =============================================================================
// 这个包负责PicoClaw的技能（Skills）管理功能。
// 技能是预定义的AI能力集合，可以扩展PicoClaw的功能。
// 技能存储在三个位置：
// - 工作区技能：~/.picoclaw/workspace/skills/
// - 全局技能：~/.picoclaw/skills/
// - 内置技能：~/.picoclaw/picoclaw/skills/
// =============================================================================

package skills

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"fmt"           // 格式化输出
	"path/filepath" // 路径操作

	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/cmd/picoclaw/internal" // 内部工具函数
	"github.com/sipeed/picoclaw/pkg/skills"            // 技能核心功能
)

// =============================================================================
// 依赖结构体 (Dependencies)
// =============================================================================
// 用于在命令间共享依赖项的结构体
type deps struct {
	workspace    string                 // 工作区路径
	installer    *skills.SkillInstaller // 技能安装器
	skillsLoader *skills.SkillsLoader   // 技能加载器
}

// =============================================================================
// NewSkillsCommand - 创建技能命令
// =============================================================================
// 这是skills模块的入口函数，创建技能相关的子命令。
// 当用户运行"picoclaw skills"时会调用此函数。
//
// 子命令说明：
// - list: 列出所有已安装的技能
// - listbuiltin: 列出内置技能
// - search: 搜索ClawHub上的技能
// - install: 安装技能
// - installbuiltin: 安装内置技能
// - remove: 移除技能
// - show: 显示技能详情
//
// 工作原理：
// - 使用PersistentPreRunE在执行任何子命令前加载配置
// - 初始化技能安装器和加载器
// - 所有子命令共享同一个依赖实例
func NewSkillsCommand() *cobra.Command {
	var d deps // 依赖实例

	// 创建技能命令
	cmd := &cobra.Command{
		Use:   "skills",
		Short: "Manage skills",
		// 在执行任何子命令前加载配置并初始化依赖
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := internal.LoadConfig()
			if err != nil {
				return fmt.Errorf("error loading config: %w", err)
			}

			// 设置工作区路径
			d.workspace = cfg.WorkspacePath()
			// 创建技能安装器
			d.installer = skills.NewSkillInstaller(d.workspace)

			// 获取全局配置目录和内置技能目录
			globalDir := filepath.Dir(internal.GetConfigPath())
			globalSkillsDir := filepath.Join(globalDir, "skills")
			builtinSkillsDir := filepath.Join(globalDir, "picoclaw", "skills")
			// 创建技能加载器
			d.skillsLoader = skills.NewSkillsLoader(d.workspace, globalSkillsDir, builtinSkillsDir)

			return nil
		},
		// 默认行为：显示帮助信息
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	installerFn := func() (*skills.SkillInstaller, error) {
		if d.installer == nil {
			return nil, fmt.Errorf("skills installer is not initialized")
		}
		return d.installer, nil
	}

	loaderFn := func() (*skills.SkillsLoader, error) {
		if d.skillsLoader == nil {
			return nil, fmt.Errorf("skills loader is not initialized")
		}
		return d.skillsLoader, nil
	}

	workspaceFn := func() (string, error) {
		if d.workspace == "" {
			return "", fmt.Errorf("workspace is not initialized")
		}
		return d.workspace, nil
	}

	cmd.AddCommand(
		newListCommand(loaderFn),
		newInstallCommand(installerFn),
		newInstallBuiltinCommand(workspaceFn),
		newListBuiltinCommand(),
		newRemoveCommand(installerFn),
		newSearchCommand(),
		newShowCommand(loaderFn),
	)

	return cmd
}
