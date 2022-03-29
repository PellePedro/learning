# Enums

```
https://exercism.org/tracks/rust/exercises/sublist

pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

Comparison::Equal
Comparison::Sublist
Comparison::Superlist
Comparison::Unequal


pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    if _first_list == _second_list {
        return Comparison::Equal;
    }
    if _first_list.len() == 0 {
        return Comparison::Sublist
    }
    if _second_list.len() == 0 {
        return Comparison::Sublist  
    }
    Comparison::Unequal
}


```
