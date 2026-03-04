// =============================================================================
// helpers.go - 定时任务辅助函数
// =============================================================================
// 这个文件包含定时任务模块的核心辅助函数，实现与定时任务服务交互的逻辑。
// =============================================================================

package cron

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"fmt"  // 格式化输出
	"time" // 时间操作

	"github.com/sipeed/picoclaw/pkg/cron" // 定时任务核心功能
)

// =============================================================================
// cronListCmd - 列出所有定时任务
// =============================================================================
// 列出所有已创建的定时任务，包括禁用的任务。
// 显示每个任务的名称、ID、调度计划、状态和下次执行时间。
//
// 参数：
//   - storePath: 定时任务存储文件路径
func cronListCmd(storePath string) {
	// 创建定时任务服务
	cs := cron.NewCronService(storePath, nil)
	// 获取所有任务（包括禁用的）
	jobs := cs.ListJobs(true)

	// 如果没有任务
	if len(jobs) == 0 {
		fmt.Println("No scheduled jobs.")
		return
	}

	// 打印任务列表
	fmt.Println("\nScheduled Jobs:")
	fmt.Println("----------------")
	for _, job := range jobs {
		// 解析调度计划
		var schedule string
		if job.Schedule.Kind == "every" && job.Schedule.EveryMS != nil {
			// 间隔任务：显示间隔秒数
			schedule = fmt.Sprintf("every %ds", *job.Schedule.EveryMS/1000)
		} else if job.Schedule.Kind == "cron" {
			// Cron任务：显示Cron表达式
			schedule = job.Schedule.Expr
		} else {
			// 一次性任务
			schedule = "one-time"
		}

		// 计算下次执行时间
		nextRun := "scheduled"
		if job.State.NextRunAtMS != nil {
			nextTime := time.UnixMilli(*job.State.NextRunAtMS)
			nextRun = nextTime.Format("2006-01-02 15:04")
		}

		// 任务状态
		status := "enabled"
		if !job.Enabled {
			status = "disabled"
		}

		// 打印任务信息
		fmt.Printf("  %s (%s)\n", job.Name, job.ID)
		fmt.Printf("    Schedule: %s\n", schedule)
		fmt.Printf("    Status: %s\n", status)
		fmt.Printf("    Next run: %s\n", nextRun)
	}
}

func cronRemoveCmd(storePath, jobID string) {
	cs := cron.NewCronService(storePath, nil)
	if cs.RemoveJob(jobID) {
		fmt.Printf("✓ Removed job %s\n", jobID)
	} else {
		fmt.Printf("✗ Job %s not found\n", jobID)
	}
}

func cronSetJobEnabled(storePath, jobID string, enabled bool) {
	cs := cron.NewCronService(storePath, nil)
	job := cs.EnableJob(jobID, enabled)
	if job != nil {
		fmt.Printf("✓ Job '%s' enabled\n", job.Name)
	} else {
		fmt.Printf("✗ Job %s not found\n", jobID)
	}
}
