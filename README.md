# 论文查重系统

## 功能描述

这是一个基于Go语言开发的论文查重系统，可以检测两篇文本文件之间的相似度。采用标准的Go项目结构，具有良好的模块化设计。

## 查重算法

本系统采用基于**编辑距离（Levenshtein距离）**的分段相似度算法：

### 核心特性
- **分段处理**：自动将长文本分割为段落和句子，提高处理长篇文档的效率
- **智能分割**：支持按段落（双换行）和句子（标点符号）进行文本分割
- **加权计算**：使用段落长度作为权重，计算综合相似度
- **最佳匹配**：每个源段落都会找到目标文本中最相似的段落进行比较

### 算法流程
1. **文本预处理**：标准化空白字符和换行符
2. **智能分段**：
   - 按双换行符分割段落
   - 长段落（>200字符）进一步按句子分割
3. **段落匹配**：为每个源段落找到最佳匹配的目标段落
4. **加权计算**：使用段落长度加权计算总体相似度

### 算法公式
- **段落相似度**：基于编辑距离计算，相似度 = 1 - (编辑距离 / 较长文本的长度)
- **总体相似度**：Σ(段落相似度 × 段落长度) / Σ(段落长度)

### 适用场景
- **长篇论文**：自动分段处理，避免整体计算的性能问题
- **多段落文档**：能够识别部分抄袭和段落重排
- **中英文混合**：完美支持Unicode字符处理

## 项目结构

```
paper-similarity-checker/
├── cmd/
│   └── main.go                    # 程序入口，处理命令行参数
├── internal/
│   ├── similarity/
│   │   ├── checker.go             # 查重核心逻辑
│   │   ├── tokenizer.go           # 文本分词处理
│   │   └── algorithms.go          # 相似度算法实现
│   └── utils/
│       ├── file_handler.go        # 文件读写处理
│       └── formatter.go           # 结果格式化
├── test/
│   ├── test_data/
│   │   ├── orig.txt               # 测试用例：原文
│   │   ├── orig_0.8_add.txt       # 测试用例：添加版
│   │   ├── orig_0.8_del.txt       # 测试用例：删减版
│   │   ├── orig_0.8_dis_1.txt     # 测试用例：微调版
│   │   ├── orig_0.8_dis_10.txt    # 测试用例：大幅调整
│   │   └── orig_0.8_dis_15.txt    # 测试用例：极大调整
│   ├── integration_test.go        # 主流程集成测试
│   └── orig_file_test.go          # 文件批量相似度测试
├── go.mod                         # Go模块文件
├── go.sum                         # 依赖锁定文件
└── README.md                      # 项目说明文档
```

## 使用方法

### 编译程序

```bash
# 编译为可执行文件
go build -o paper-checker.exe ./cmd

# 或者直接运行
go run ./cmd <原文文件路径> <抄袭版文件路径> <输出文件路径>
```


### 运行单元测试

```bash
# 运行所有测试（推荐）
## 测试

# 运行某个模块的测试
go test ./internal/similarity
go test ./internal/utils

# 运行带详细输出的测试
go test -v ./internal/...

# 运行特定测试函数
go test -v ./internal/similarity -run TestCalculateSimilarity
go test -v ./internal/utils -run TestFormatSimilarityAsPercentage
```

各模块均有对应的测试文件（如 checker_test.go、formatter_test.go、file_handler_test.go），可直接运行。

### 运行单元测试

```bash
# 运行所有测试
go test ./test/

# 运行带详细输出的测试
go test -v ./test/

# 运行特定测试
go test -v ./test/ -run TestSimilarityChecker
```

## 模块说明

### 核心模块

- **similarity.Checker**: 查重检查器，协调各个算法组件
- **similarity.Tokenizer**: 文本分词器，处理文本预处理和分词
- **similarity.Algorithms**: 算法库，实现各种相似度计算算法

### 工具模块

- **utils.FileHandler**: 文件处理器，负责文件读写操作
- **utils.Formatter**: 格式化器，负责结果格式化输出

## 输出格式

程序将相似度结果（0.00-1.00）与比对文本路径保存到指定的输出文件中，保留两位小数。

## 开发说明

本项目遵循Go语言的标准项目布局：

- `cmd/`: 应用程序入口
- `internal/`: 私有应用程序代码
- `test/`: 测试代码和数据
- 使用Go Modules进行依赖管理