## Buffer Management

disk にあるデータの仮想的な配置が block である。DB からデータを取り出す時、block を page と呼ばれるメモリ空間におく必要がある。
**buffer manager** は、データを保持する page を管理する責任を持つ。
buffer manager は page を *pin* / *unpin* することで、page の二つのステータスをコントロールする。
**Buffer** は、page と配置された block の識別子、page が *pin* されているかどうか、といった情報を持っているものである。**Buffer pool** はその集合である。

以下のプロコトルを経て、client は block にアクセスする。

1. clinet が buffer pool にある page を pin することを buffer manager に要求する。
2. client は自由にその page のデータにアクセスする。
3. client がその page へのアクセスを終了させると、buffer manager に その page を *unpin* することを要求する。


client が buffer manager に pin することを要求するとき、次の 4 つの可能性の内のいずれかになる。

- block の内容が、buffer のどこかの page にある。
  - その page が pin されている。
  - その page が unpin されている。
- block の内容が、いずれの buffer にも存在しない。
  - buffer pool に unpin されている page が少なくとも一つはある。
  - buffer pool 内の全ての page が pin されている。


1 の場合を考える。page は複数の client に pin されることができるので、追加で pin をして、client に buffer manager はその page を返却するだけである。
2 の場合を考える。これはその page を使っていた clinet がその page へのアクセスを終了した時の状態である。まだ、block のデータは buffer page にあるので、シンプルにその page を pin して、page を client に
返却すれば良い。
3 の場合を考える。これは、buffer manager が、disk から page に block を読みこむ必要がある。まず、buffer manager は、unpin されている page を一つ選択する。もし、この選択された page が変更されていたら、
その内容を disk に書き込む。これを **flush** と呼ぶ。その後、page に block の内容を読み込んで、その page を pin する。
4 の場合を考える。これは、利用可能なメモリがない状態で、利用可能になるまで、client を wait list に入れる。

