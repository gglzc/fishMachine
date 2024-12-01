# gRPC Fish Machine Server

## 介紹

`gRPC Fish Machine Server` 是一個基於 gRPC 的捕魚遊戲服務器，專為高效通信設計。它支持多人實時交互、魚群生成和多種遊戲道具，適用於各種捕魚類遊戲的後端支持。

---

## 功能特性

- **實時多人支持**：高效的 gRPC 通信，保障多用戶流暢交互。
- **魚群生成**：支持多種魚類類型與隨機生成規則。
- **道具效果**：內建多種增強型遊戲道具，提升玩家體驗。
- **高擴展性**：模塊化設計，便於功能擴展和維護。
- **高效性能**：基於 gRPC 的高性能低延遲通信框架。

---

## 系統需求


- **依賴項目**：
  - gRPC 庫
  - Protobuf 編譯器
  - 數據庫（MySQL）

---

## 安裝與運行

### 1. 克隆項目
```bash
git clone https://github.com/gglzc/fishMachine.git
cd fish-machine-server
