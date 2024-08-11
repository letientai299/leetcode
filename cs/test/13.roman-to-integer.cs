using NUnit.Framework;

namespace Test;

public partial class SolutionTest
{
  [TestCase("I", 1)]
  [TestCase("II", 2)]
  [TestCase("III", 3)]
  [TestCase("IV", 4)]
  [TestCase("V", 5)]
  [TestCase("VI", 6)]
  [TestCase("MCMXCIV", 1994)]
  public void TestRomanToInt(string s, int expected)
  {
    var got = new Solution().RomanToInt(s);
    Assert.That(got, Is.EqualTo(expected));
  }
}
