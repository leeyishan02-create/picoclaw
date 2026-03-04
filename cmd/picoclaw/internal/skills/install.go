// =============================================================================
// install.go - 安装技能命令实现
// =============================================================================
// 这个文件实现了PicoClaw的技能安装功能。
// 用户可以从GitHub安装技能。
// =============================================================================

package skills

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"fmt" // 格式化输出

	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/cmd/picoclaw/internal" // 内部工具函数
	"github.com/sipeed/picoclaw/pkg/skills"            // 技能核心功能
)

// =============================================================================
// newInstallCommand - 创建安装命令
// =============================================================================
// 创建"picoclaw skills install"子命令。
// 从GitHub安装技能。
//
// 参数：
//   - installerFn: 返回技能安装器的函数
//
// 命令行参数：
//   - --registry: 指定技能仓库（可选）
//
// 位置参数：
//   - args[0]: GitHub仓库地址（如 sipeed/picoclaw-skills/weather）
//   - args[1]: 当使用--registry时，技能slug
//
// 使用示例：
//
//	# 从GitHub安装技能
//	picoclaw skills install sipeed/picoclaw-skills/weather
//
//	# 从ClawHub安装
//	picoclaw skills install --registry clawhub github
func newInstallCommand(installerFn func() (*skills.SkillInstaller, error)) *cobra.Command {
	var registry string // 技能仓库

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install skill from GitHub",
		Example: `
picoclaw skills install sipeed/picoclaw-skills/weather
picoclaw skills install --registry clawhub github
`,
		// 参数验证函数
		Args: func(cmd *cobra.Command, args []string) error {
			if registry != "" {
				// 如果指定了registry，需要两个参数
				if len(args) != 2 {
					return fmt.Errorf("when --registry is set, exactly 2 arguments are required: <name> <slug>")
				}
				return nil
			}

			// 默认需要一个参数
			if len(args) != 1 {
				return fmt.Errorf("exactly 1 argument is required: <github>")
			}

			return nil
		},
		RunE: func(_ *cobra.Command, args []string) error {
			installer, err := installerFn()
			if err != nil {
				return err
			}

			if registry != "" {
				cfg, err := internal.LoadConfig()
				if err != nil {
					return err
				}

				return skillsInstallFromRegistry(cfg, args[0], args[1])
			}

			return skillsInstallCmd(installer, args[0])
		},
	}

	cmd.Flags().StringVar(&registry, "registry", "", "Install from registry: --registry <name> <slug>")

	return cmd
}
