using System;
using System.Collections.Generic;
using NUnit.Framework;

public partial class SolutionTest
{
    [Test]
    public void TestReverse()
    {
        var sln = new Solution();
        var tests = new List<Tuple<int, int>>
        {
            new Tuple<int, int>(123, 321),
            new Tuple<int, int>(-123, -321),
            new Tuple<int, int>(0, 0),
            new Tuple<int, int>(1234567899, 0),
        };

        foreach (var ( input, expected) in tests)
        {
            var got = sln.Reverse(input);
            Assert.AreEqual(got, expected);
        }
    }
}