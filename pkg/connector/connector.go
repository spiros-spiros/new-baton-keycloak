package connector

import (
	"context"
	"io"

	"github.com/conductorone/baton-keycloak/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
)

const accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ4MFdpUXNYUTBmWk5qWXJaX0xDSmF2a3pCZE1LTWxIbmtublljYmNxeDFJIn0.eyJleHAiOjE2OTkzNTQ5MTUsImlhdCI6MTY5OTMxODkxNSwianRpIjoiMDA2ZGJjY2ItNjkxYy00ZjgyLTllYzctODhlYzAyMTU4YzAxIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tYXN0ZXIiLCJzdWIiOiI5MGJjYWQ1NS00ODExLTRmYzktYTY3ZS00ZDEyNDU2OTE0ZmMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhZG1pbi1jbGkiLCJzZXNzaW9uX3N0YXRlIjoiNWE3N2E1ODEtYjljNy00NDgwLThmNWMtYjM2NTNiM2IxMjcyIiwiYWNyIjoiMSIsInNjb3BlIjoiZW1haWwgcHJvZmlsZSIsInNpZCI6IjVhNzdhNTgxLWI5YzctNDQ4MC04ZjVjLWIzNjUzYjNiMTI3MiIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoia2VzaGF2In0.XF567H0OzYKv08zlihJ7Dg9DHjhiURMV9DQUM5gzofYvAiCl6YtVcqWMjMvbXdUpwR4AFfRdxfrWVORMEYYqD2yv5JL8mX3VspaKMD-Ypy2mk7sqR6LQgJiyMm9nOMvV_7vzqa8Yy7upXzyiVOQcPWKBqcQuaGeQPGHPSs222j6HhuCYqmbVmza_k-Qu4lnCumxRbCPrdMbAiiTGXbHHh8wcIxUBpj65ekXC_qQEnXhdGT0F98yZDUMdCHsEPsjnoUHFSp2_SCJ0FOqzVr5IEQ2TVNnGpW_sWRA2x6qbOv78j03dlmTbovCqzC6ZiH5MP6FEmv5pPflFzkSVxDo60Q"

type Connector struct{ keyCloakClient *client.Client }

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (d *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(d.keyCloakClient),
		newGroupBuilder(d.keyCloakClient),
	}
}

// Asset takes an input AssetRef and attempts to fetch it using the connector's authenticated http client
// It streams a response, always starting with a metadata object, following by chunked payloads for the asset.
func (d *Connector) Asset(ctx context.Context, asset *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (d *Connector) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "Keycloak",
		Description: "A baton connector for the keycloak idP",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (d *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	return nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context, baseUrl string) (*Connector, error) {
	return &Connector{keyCloakClient: client.New(accessToken, baseUrl)}, nil
}
