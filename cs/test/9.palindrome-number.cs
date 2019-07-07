using System;
using System.Collections.Generic;
using NUnit.Framework;

public partial class SolutionTest
{
    [Test]
    public void TestIsPalindrome()
    {
        var sln = new Solution();
        var tests = new List<Tuple<int, bool>>
        {
            new Tuple<int, bool>(-1, false),
            new Tuple<int, bool>(12, false),
            new Tuple<int, bool>(-121, false),
            new Tuple<int, bool>(121, true),
            new Tuple<int, bool>(0, true),
            new Tuple<int, bool>(1, true),
            new Tuple<int, bool>(11, true),
            new Tuple<int, bool>(101, true),
            new Tuple<int, bool>(1001, true),
        };


        foreach (var ( input, expected) in tests)
        {
            var got = sln.IsPalindrome(input);
        }
    }
}