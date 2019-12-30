<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../find-the-team-size "Find the Team Size")
　　　　　　　　　　　　　　　　
[Next >](../all-elements-in-two-binary-search-trees "All Elements in Two Binary Search Trees")

## [1304. Find N Unique Integers Sum up to Zero (Easy)](https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero "和为零的N个唯一整数")

<p>Given an integer <code>n</code>, return <strong>any</strong> array containing <code>n</code> <strong>unique</strong>&nbsp;integers such that they add up to 0.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 5
<strong>Output:</strong> [-7,-1,1,3,4]
<strong>Explanation:</strong> These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 3
<strong>Output:</strong> [-1,0,1]
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> n = 1
<strong>Output:</strong> [0]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 1000</code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Return an array where the values are symmetric. (+x , -x).
</details>

<details>
<summary>Hint 2</summary>
If n is odd, append value 0 in your returned array.
</details>