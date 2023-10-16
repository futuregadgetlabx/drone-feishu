package template

const (
	PUSH_SUCCESS = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**构建时间**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**仓库地址**\n[{{ .Repo.Owner }}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 Built by [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n🔀 构建分支: [{{ .Build.Branch }}]({{ .Repo.Url }}/tree/{{ .Build.Branch }})\n✅ Git Commit: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 Commit message: {{ .Build.CommitMessage }}\n🛠️ 构建任务: [#{{ .Build.Number }}]({{ .Build.Link }})\n⏱️ 构建耗时: {{ .Build.CostTime }}s"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "部署上线"
          },
          "type": "primary",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "我已知悉"
          },
          "type": "default"
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所](https://github.com/futuregadgetlabx)",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "green",
    "title": {
      "content": "🎉【Drone CI】代码编译成功",
      "tag": "plain_text"
    }
  }
}
`

	PR_SUCCESS = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**构建时间**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**仓库地址**\n[{{ .Repo.Owner }}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 **Built by** [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n📌 **Pull Request**: ***{{ .Build.SourceBranch }} -> {{ .Build.TargetBranch }}***\n{{ .Build.PullRequestTitle }}\n🖇️ **Commit**: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 **Commit message**: {{ .Build.CommitMessage }}\n🛠️ **构建任务**: [#{{ .Build.Number }}]({{ .Build.Link }})\n⏱️ **构建耗时**: {{ .Build.CostTime }}s"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Merge PR"
          },
          "type": "primary",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Code Review"
          },
          "type": "default",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Close PR"
          },
          "type": "danger",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所](https://github.com/futuregadgetlabx)",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "green",
    "title": {
      "content": "🎉【Drone CI】编译通过",
      "tag": "plain_text"
    }
  }
}
`

	PUSH_FAILURE = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**构建时间**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**仓库地址**\n[{{ .Repo.Owner }}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 Built by [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n🔀 构建分支: [{{ .Build.Branch }}]({{ .Repo.Url }}/tree/{{ .Build.Branch }})\n✅ Git Commit: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 Commit message: {{ .Build.CommitMessage }}\n🛠️ 构建任务: [#{{ .Build.Number }}]({{ .Build.Link }})\n❌ Failed stages: {{ .Build.FailedStages }}\n🔥️ Failed steps: {{ .Build.FailedSteps }}"
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所](https://github.com/futuregadgetlabx)",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "red",
    "title": {
      "content": "🚒【Drone CI】代码编译失败",
      "tag": "plain_text"
    }
  }
}
`

	PR_FAILURE = `
`
)
