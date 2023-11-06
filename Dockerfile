FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-keycloak"]
COPY baton-keycloak /