# Sugoroku_and_Pieces
## atcoder:https://www.ioi-jp.org/joi/2018/2019-yo/2019-yo-t2.html
###問題文
JOI 君はすごろくを持っている．このすごろくは 2019 個のマスが横一列に並んだ形をしている．これらのマスには，左端のスタートマスから右端のゴールマスへと順に 1 から 2019 までの番号がついている．

現在このすごろくの上には，N 個の駒が置かれている．これらの駒には，スタートに近い順に 1 から N までの番号がついている．駒 i (1 ≦ i ≦ N) は，マス X_i に置かれている．すべての駒は異なるマスに置かれている．

JOI 君はこれから M 回の操作を行う．j 回目 (1 ≦ j ≦ M) の操作では，駒 A_j を 1 マス先へ進める．ただし，移動元のマスがゴールマスであった場合，もしくは移動先のマスに別の駒が置かれている場合，駒 A_j は進まず，位置は変わらない．

すべての操作が終了した時点で，各駒が置かれているマスを求めよ．
