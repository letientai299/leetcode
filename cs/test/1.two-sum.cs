using System;
using System.Collections.Generic;
using NUnit.Framework;

public partial class SolutionTest
{
    [Test]
    public void TestTwoSum()
    {
        var sln = new Solution();
        var tests = new List<Tuple<int[], int, int[]>>
        {
            new Tuple<int[], int, int[]>(new[] {1, 2, 3}, 4, new[] {0, 2}),
        };

        foreach (var (nums, k, expected) in tests)
        {
            var got = sln.TwoSum(nums, k);
            Assert.That(got, Is.EquivalentTo(expected));
        }
    }
}