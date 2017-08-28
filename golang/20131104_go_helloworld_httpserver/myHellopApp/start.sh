# mvn archetype:generate -DgroupId=com.example -DartifactId=myHellopApp -DinteractiveMode=false
# cd myHellopApp

mvn compile
mvn exec:java -Dexec.mainClass="com.example.App"

