using System;
using System.Collections.Generic;
using NUnit.Framework;

public partial class SolutionTest
{
    [Test]
    public void TestRomanToInt()
    {
        var sln = new Solution();
        var tests = new List<Tuple<string, int>>
        {
            new Tuple<string, int>("I", 1),
            new Tuple<string, int>("II", 2),
            new Tuple<string, int>("III", 3),
            new Tuple<string, int>("IV", 4),
            new Tuple<string, int>("V", 5),
            new Tuple<string, int>("VI", 6),
            new Tuple<string, int>("MCMXCIV", 1994),
        };

        foreach (var (input, expected) in tests)
        {
            var got = sln.RomanToInt(input);
            Assert.AreEqual(expected, got);
        }
    }
}