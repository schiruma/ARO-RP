package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder provides operations to manage the artifact property of the microsoft.graph.security.indicator entity.
type ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetQueryParameters the artifact related to this indicator.
type ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetQueryParameters
}
// NewThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderInternal instantiates a new ArtifactRequestBuilder and sets the default values.
func NewThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder) {
    m := &ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/intelligenceProfileIndicators/{intelligenceProfileIndicator%2Did}/artifact{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder instantiates a new ArtifactRequestBuilder and sets the default values.
func NewThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the artifact related to this indicator.
func (m *ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Artifactable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Artifactable), nil
}
// ToGetRequestInformation the artifact related to this indicator.
func (m *ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceIntelligenceProfileIndicatorsItemArtifactRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
