struct Solution {}

impl Solution {
    pub fn successful_pairs(spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
        potions.sort();
        println!("{:?}", potions);
        return spells
            .into_iter()
            .map(|spell| -> i32 {
                let position = potions.partition_point(|&input| {
                    let after_spell = (input as i64) * (spell as i64);
                    return after_spell < success;
                });
                return (potions.len() - position) as i32;
            })
            .collect();
    }
}
/*
72, 72, 71, 68, 72, 71, 29, 71, 72, 72, 72, 68, 68, 72, 59, 71, 71, 68, 71, 72, 68, 72, 71, 59, 72, 72, 71, 72, 72, 72, 51, 71, 0, 71, 59, 72, 72, 72, 68, 71, 0, 72, 71, 71, 71, 72, 68, 71, 45, 71, 71, 68, 71, 71, 59, 71, 29, 72, 68, 45, 72, 71, 68, 71, 45, 68, 72, 51, 59, 71, 71, 68, 68, 71, 0
*/

#[cfg(test)]
mod test {
    use crate::Solution::Solution;

    #[test]
    fn test() {
        assert_eq!(
            vec![69],
            Solution::successful_pairs(
                vec![22],
                vec![
                    30, 11, 5, 20, 19, 36, 39, 24, 20, 37, 33, 22, 32, 28, 36, 24, 40, 27, 36, 37,
                    38, 23, 39, 11, 40, 19, 37, 32, 25, 29, 28, 37, 31, 36, 32, 40, 38, 22, 17, 38,
                    20, 33, 29, 17, 36, 33, 35, 25, 28, 18, 17, 19, 40, 27, 40, 28, 40, 40, 40, 39,
                    17, 34, 36, 11, 22, 29, 22, 35, 35, 22, 18, 34
                ],
                135,
            )
        );
        assert_eq!(
            vec![2, 0, 2],
            Solution::successful_pairs(vec![3, 1, 2], vec![8, 5, 8], 16,)
        );
    }
}
