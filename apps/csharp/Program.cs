using Amazon.CDK;

namespace Main
{
  public class Program
  {
    static void Main(string[] args)
    {
      var app = new App(null);
      new Stack(app, "TestStack", null);
      app.Synth();
    }
  }
}
