## Concurrency Manegement

```
T1 : W(b1); W(b2)
T2 : W(b1); W(b2)
```

transaction T1, T2 が直列で実行されるとき、_serial schedule_ と呼ぶ。non-serial schedule であるが、ある serial schedule と同じ結果をもたらすものを _serializable_ と呼ぶ。

```
W1(b1); W2(b1); W1(b2); W2(b2)
```

これは、T1 を実行した後に、T2 を実行する serial schedule と同じ結果になるので、seralizable である。一方で、次の schedule を考える。

```
W1(b1); W2(b1); W2(b2); W1(b1)
```

これは、T1 が文字 'X', T2 が文字 'Y' を書き込んでいるとすると、最終的な結果は、block1 には、'X' が、block2 には、'Y' が書き込まれ、直列実行した時と必ず結果が異なる。
そのため、この schedule は _non-serializable_ である。
ACID property の 独立性の性質は、他のトランザクションに影響を受けないことなので、non-serializable の時、この性質は満たされない。

### Lock Table

データベースエンジンは、全ての schedule が serializable になることを保証する責任を持つ。そのための一般的な技術が **lock** である。
各 block は、2種類の lock を有する。それは、共有 lock ( slock ) と排他 lock ( xlock ) である。ある transaction が、xlock をある block にかけると、他の transaction は
いかなる lock もその block に対してかけることができない。一方で、ある transaction が slock をかけた場合は、他の tansaction が slock をかける事は可能であるが、xlock をかけることはできない。
