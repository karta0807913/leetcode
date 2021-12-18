#include <algorithm>
#include <cassert>
#include <iostream>
#include <istream>
#include <vector>

using namespace std;

class Solution {
public:
  void sort(vector<int> &nums, int start, int end) {
    while (nums[start] < nums[end] && start != end) {
      end -= 1;
    }
    while (start != end) {
      swap(nums[start], nums[end]);
      end -= 1;
    }
  }

  int search(vector<int> &nums, int start, int end, int target) {
    return this->binary_search(nums, start, end, target);
  }

  int normal_search(vector<int> &nums, int start, int end, int target) {
    int index = end;
    for (; start <= index; --index) {
      if (nums[index] <= target) {
        break;
      }
    }
    if (index < start || index + 1 > end) {
      return -1;
    }
    // for(; start <= end; ++start) {
    //     cout << nums[start] << endl;
    //     if(nums[start] != current) {
    //         return start - 1;
    //     }
    // }
    return index + 1;
  }

  int binary_search(vector<int> &nums, int start, int end, int target) {
    int origin = end;
    // cout << "start " << start << " " << end << ", target " << target << endl;
    while (start <= end) {
      int current_index = (end + start) / 2;
      if (nums[current_index] < target) {
        start = current_index + 1;
      } else if (nums[current_index] > target) {
        end = current_index - 1;
      } else {
        start = current_index + 1;
        end = max(start, end);
        // cout << "equal " << start << " " << end << endl;
        while (start <= origin) {
          if (nums[start] > target) {
            return start;
          }
          start += 1;
        }
        return -1;
      }
    }
    // cout << nums[end] << " END " << end << endl;
    if (nums[end] > target) {
      return end;
    }
    if (nums[start] > target) {
      return start;
    }
    return -1;
  }

  void nextPermutation(vector<int> &nums) {
    if (nums.size() <= 1) {
      return;
    }
    auto start = nums.size() - 2, end = nums.size() - 1;
    while (start != -1) {
      // cout << "target " << nums[start] << " start " << nums[start + 1]
      //      << " end " << nums[end] << endl;
      // for (auto iter = nums.begin(); iter != nums.end(); ++iter) {
      //   cout << *iter << " ";
      // }
      // cout << endl;
      if (nums[start] < nums[end]) {
        int index = this->search(nums, start + 1, end, nums[start]);
        if (index > -1) {
          swap(nums[index], nums[start]);
          return;
        }
        // cout << "do " << index << endl;
      }
      this->sort(nums, start, end);
      // for (auto iter = nums.begin(); iter != nums.end(); ++iter) {
      //   cout << *iter << " ";
      // }
      // cout << endl;
      start -= 1;
    }
  }
};

int main() {
  Solution solution;
  vector<int> test = {5, 1, 2, 3, 4};
  int result = solution.search(test, 1, test.size() - 1, 3);
  cout << result << endl;
  assert(test[result] == 4);
  solution.sort(test, 0, test.size() - 1);
  for (auto iter = test.begin(); iter != test.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  test = {0,0,0,1,2,4};
  result = solution.search(test, 2, test.size() - 1, 0);
  assert(test[result] == 1);

  test = {1, 1, 5};
  result = solution.search(test, 1, test.size() - 1, 1);
  assert(test[result] == 5);

  test = {1, 5, 1};
  solution.nextPermutation(test);
  for (auto iter = test.begin(); iter != test.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;
  assert(test[0] == 5);
  assert(test[1] == 1);
  assert(test[2] == 1);

  test = {2, 1, 3};
  result = solution.search(test, 1, test.size() - 1, 2);
  assert(test[result] == 3);

  test = {1, 2, 3};
  result = solution.search(test, 1, test.size() - 1, 1);
  assert(test[result] == 2);
  solution.nextPermutation(test);
  assert(test[0] == 1);
  assert(test[1] == 3);
  assert(test[2] == 2);

  test = {2, 2, 1, 2, 2, 3, 4, 5, 7};
  result = solution.search(test, 2, test.size() - 1, 2);
  cout << result << endl;
  assert(test[result] == 3);

  vector<int> input = {1, 2, 3};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  input = {3, 2, 1};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  input = {1};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  input = {1, 1, 3, 2};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  input = {2, 1, 3};
  assert(input[solution.search(input, 1, 2, 2)] == 3);
  input = {2, 3, 1};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  input = {2, 3, 1, 1};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;

  input = {5, 4, 7, 5, 3, 2};
  solution.nextPermutation(input);
  for (auto iter = input.begin(); iter != input.end(); ++iter) {
    cout << *iter << " ";
  }
  cout << endl;
}
