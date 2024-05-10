struct Solution {}

impl Solution {
    pub fn min_operations(mut nums: Vec<i32>, queries: Vec<i32>) -> Vec<i64> {
        let mut previous: i64 = 0;
        nums.sort();
        nums.insert(0, 0);
        let prefix_sum: Vec<i64> = nums
            .iter()
            .map(|num| -> i64 {
                previous += *num as i64;
                return previous;
            })
            .collect();
        return queries
            .into_iter()
            .map(|query| -> i64 {
                let bigger_or_equal = nums.binary_search(&query).unwrap_or_else(|x| x) - 1;
                return (query as i64 * bigger_or_equal as i64 - prefix_sum[bigger_or_equal] * 2)
                    + (prefix_sum.last().unwrap()
                        - ((nums.len() - 1 - bigger_or_equal) as i64 * query as i64));
            })
            .collect();
    }
}

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn test_min_operations() {
        assert_eq!(
            vec![14, 10],
            Solution::min_operations(vec![3, 1, 6, 8], vec![1, 5],)
        )
    }
}
