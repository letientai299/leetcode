using NUnit.Framework;

namespace Test;

public partial class SolutionTest
{
  [TestCase(-1, false)]
  [TestCase(12, false)]
  [TestCase(-121, false)]
  [TestCase(121, true)]
  [TestCase(0, true)]
  [TestCase(1, true)]
  [TestCase(11, true)]
  [TestCase(101, true)]
  [TestCase(1001, true)]
  public void TestIsPalindrome(int n, bool want)
  {
    var got = new Solution().IsPalindrome(n);
    Assert.That(got, Is.EqualTo(want));
  }
}
