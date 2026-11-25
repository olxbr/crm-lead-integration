FROM eclipse-temurin:17
RUN groupadd --system app && useradd --system --gid app --no-create-home app
ARG JAR_FILE
COPY ${JAR_FILE} app.jar
USER app
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
  CMD curl -fs http://localhost:8080/actuator/health || exit 1
ENTRYPOINT ["java","-jar","/app.jar"]
