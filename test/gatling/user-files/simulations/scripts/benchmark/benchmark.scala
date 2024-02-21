import scala.concurrent.duration._

import scala.util.Random

import io.gatling.core.Predef._
import io.gatling.http.Predef._


class ApiGatewayBenchmarkSimulation
  extends Simulation {

  val httpProtocol = http
    .baseUrl("http://localhost:8080")
    .userAgentHeader("Fake Api Simulation")

  val testPath1 = scenario("Path 1")
    .feed(tsv("api-without-delay.tsv").circular())
    .exec(
      http("test")
      .post("/servicea").body(StringBody("#{payload}"))
      .header("content-type", "application/json")
      // 201 pros casos de sucesso :)
      // 422 pra requests inválidos :|
      // 400 pra requests bosta tipo data errada, tipos errados, etc. :(
      .check(status.in(200, 201, 422, 400))
    )

  val testPath2 = scenario("Path 2")
  .feed(tsv("api-without-delay.tsv").circular())
    .exec(
      http("test")
      .post("/serviceb").body(StringBody("#{payload}"))
      .header("content-type", "application/json")
      // 201 pros casos de sucesso :)
      // 422 pra requests inválidos :|
      // 400 pra requests bosta tipo data errada, tipos errados, etc. :(
      .check(status.in(200, 201, 422, 400))
    )

  val testPath3 = scenario("Path 3")
  .feed(tsv("api-without-delay.tsv").circular())
    .exec(
      http("test")
      .post("/servicec").body(StringBody("#{payload}"))
      .header("content-type", "application/json")
      // 201 pros casos de sucesso :)
      // 422 pra requests inválidos :|
      // 400 pra requests bosta tipo data errada, tipos errados, etc. :(
      .check(status.in(200, 201, 422, 400))
    )

  setUp(
    testPath1.inject(
      constantUsersPerSec(2).during(10.seconds), // warm up
      constantUsersPerSec(5).during(15.seconds).randomized, // are you ready?

      rampUsersPerSec(6).to(1000).during(3.minutes) // lezzz go!!!
    ),
    testPath2.inject(
      constantUsersPerSec(2).during(10.seconds), // warm up
      constantUsersPerSec(5).during(15.seconds).randomized, // are you ready?

      rampUsersPerSec(6).to(1000).during(3.minutes) // lezzz go!!!
    ),
    testPath3.inject(
      constantUsersPerSec(2).during(10.seconds), // warm up
      constantUsersPerSec(5).during(15.seconds).randomized, // are you ready?

      rampUsersPerSec(6).to(1000).during(3.minutes) // lezzz go!!!
    )
  ).protocols(httpProtocol)
}
