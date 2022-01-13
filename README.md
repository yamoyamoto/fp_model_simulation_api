# 連想記憶シュミレーションのためのミニAPIサーバー

## Overview

物理学科の卒業研究テーマ「*ホップフィールド模型と連想記憶の関係*」で使用する、簡単な連想記憶のシュミレーションアプリ。(兼golangの練習)
ここではアプリケーションに必要なAPIサーバーを実装している。

また、フロントのコードは以下のリポジトリで管理している。

<https://github.com/yamoyamoto/fp_model_simulation>

## 連想記憶

N個の1または-1の値を取る素子を考え、以下の式でダイナミクスを定義する。

<img src=
"https://render.githubusercontent.com/render/math?math=%5Cdisplaystyle+%5Cbegin%7Balign%2A%7D%0AS_i%28t%2B%5CDelta+t%29+%3D+sgn%28%5Csum_%7B%5Cj%3D1%7D%5EN+J_%7Bij%7DS_j%29%0A%5Cend%7Balign%2A%7D%0A"
alt="\begin{align*}
S_i(t+\Delta t) = sgn(\sum_{\j=1}^N J_{ij}S_j)
\end{align*}
">

また、Jは結合定数と呼ばれ、パターンを記憶する役割を担う。
Jを以下の式で表すことにする。

<img src=
"https://render.githubusercontent.com/render/math?math=%5Cdisplaystyle+%5Cbegin%7Balign%2A%7D%0AJ_%7Bij%7D+%3D+%5Clambda%5Csum_%7B%5Cmu%3D1%7D%5Ep%5Cxi_i%5E%5Cmu%5Cxi_j%5E%5Cmu%0A%5Cend%7Balign%2A%7D%0A"
alt="\begin{align*}
J_{ij} = \lambda\sum_{\mu=1}^p\xi_i^\mu\xi_j^\mu
\end{align*}
">

これを「ヘッブの法則」という。また係数λは、学習の振幅や学習率と呼ばれる。
ξは「μ番目の記銘パターンにおけるi番目の素子の値」を示し、pは記名するパターン数を示す。

ヘッブの法則で表された結合定数を使ってダイナミクスを実行することを考える。

ダイナミクスの実行中、記名したパターンにひとたび遷移すると、それ以降はいくらダイナミクスを実行しても別のパターンには遷移しなくなる。
(上の式を用いて証明可能)

このヘッブの法則から導かれた結合定数の性質を使えば「記憶の想起」を実現できる。
