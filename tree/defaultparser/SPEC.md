# Specification of defaultparser

## Summary
```
<CompleteQuery>  ::= <Expression>
<Expression>   	 ::= <Term>{ "OR" <Term>}
<Term>      	 ::= <Factor> "AND" <Term> | <Factor>
<Factor>         ::= <String> | "(" <Expression> ")" | "-"<Factor>
<String>         ::= <Key>:<Value>
```

## Examples

### Valid queries

```
key1:value1 AND key2:value2
key1:value1 OR key2:value2 OR key3:value3 OR key4:value4
key1:value1 AND (key2:value2 OR key3:value3)
key1:value1 AND -key2:value2
-(key1:value1 OR key2:value2)
key1:>= 1
```


### Invalid queries

```
key1:value1 AND
) key2:value2
(key1:value1
key1:value1 OR key2:value2)
key1 >= 1
```

