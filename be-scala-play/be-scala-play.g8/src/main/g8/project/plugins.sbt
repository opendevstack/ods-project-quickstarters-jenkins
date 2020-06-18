credentials += Credentials(Path.userHome / ".sbt" / ".credentials")

addSbtPlugin("com.typesafe.play"        % "sbt-plugin"          % "2.8.2")
addSbtPlugin("org.scalameta"            % "sbt-scalafmt"        % "2.4.0")
addSbtPlugin("org.foundweekends.giter8" % "sbt-giter8-scaffold" % "0.11.0")
