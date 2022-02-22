package com.github.romainmuller.cdkimportperf;

import software.amazon.awscdk.App;
import software.amazon.awscdk.Stack;

public final class Main {
  public static void main(final String[] args) {
    final App app = new App();
    new Stack(app, "TestStack");
    app.synth();
  }
}
