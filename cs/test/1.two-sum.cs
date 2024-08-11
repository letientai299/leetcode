using NUnit.Framework;

namespace Test;

public partial class SolutionTest
{
  [TestCase(new[] { 1, 2, 3 }, 4, new[] { 0, 2 })]
  [TestCase(new[] { 2, 7, 11, 15 }, 9, new[] { 0, 1 })]
  public void TestTwoSum(int[] nums, int k, int[] expected)
  {
    var got = new Solution().TwoSum(nums, k);
    Assert.That(got, Is.EquivalentTo(expected));
  }
}
