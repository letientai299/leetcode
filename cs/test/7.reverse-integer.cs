using NUnit.Framework;

namespace Test;

public partial class SolutionTest
{
  [TestCase(123, 321)]
  [TestCase(-123, -321)]
  [TestCase(0, 0)]
  [TestCase(1234567899, 0)]
  public void TestReverse(int n, int expected)
  {
    var got = new Solution().Reverse(n);
    Assert.That(got, Is.EqualTo(expected));
  }
}
