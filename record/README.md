# Record Management

## Spanned Versus Unspanned Records
Record の block への配置方法として、1 つの record が複数の block を跨がないで配置する方法と block に入りきれなかった分を別の block に配置する、つまり複数の block を跨いで配置する方法が考えられる。
record を分割して保存するときの、record を spanned record という。これらの、trade-off を考える。
- unspanned record のデメリットは、充填効率の悪さにある。
- spanned record のデメリットは、record access の複雑さにある。

file が全て同じテーブルの record で構成されているとき、*homegeneous* という。この trade-off について考える。

- single table SQL の時は、*homogeneous* は効率的であるが、JOIN などの multi-table queries の時は効率が悪くなる。 それは、*nonhomogeneous* なら、同じ ( または近く ) block に、保存することで JOIN 時の disk access を減らせるためである。

## Fixed-Length Versus Variable-Length Fields

全ての table のフィールドは型情報を持っている。record manager は各フィールドが *fixed-length* または *variable-length* のどちらを選択する必要がある。integer や floating-point numbers は 4 byte binary として保存できる。しかし、string は variable-length になる。
variable-length には高度な複雑さがある。例えば、block の中間にある record を変更しようとする時、variable-length で大きい値に変更する時、record を再配置しなおす必要がある。また、あまりに大きいと、別の block へ移動させる必要もあるだろう。

したがって、record manager は可能な限り、fixed-lengh を選択しようとするが、string filed の異なる 3 つの表現から選択することができる。

1. variable-length representation で、string の正確なサイズ分だけ block に割り当てる。
2. fixed length representation で、record の外側に strnig を保存し、その場所への固定長の参照を配置する。
3. fixed length representation で、そのサイズに関わらず、固定長の size を割り当てる。

標準的な SQL は、3 種類の string datatype を与える: char, varchar, clob. char(n) 型は、正確に n 文字の string である。varchar(n) と clob(n) は最大で n 文字の string である。varchar と clob では n の予測される大きさが異なる。varchar は小さい値で、4K を超えないが、clob はギガ単位である。
char 型の時は、上の 3 を選択すると最も効率が良い。varchar(n) の時は、上の 1 を選択すると良い。それは、n が大きくないので、record をそれほど大きく変化させないためである。n がとても小さいと、3 を選択することがある。clob は 3 を選択すると良い。record を小さくすることができ、より管理しやすくなるためである。


