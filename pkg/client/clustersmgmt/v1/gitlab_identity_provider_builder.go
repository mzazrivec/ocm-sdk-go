/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// GitlabIdentityProviderBuilder contains the data and logic needed to build 'gitlab_identity_provider' objects.
//
// Details for `gitlab` identity providers.
type GitlabIdentityProviderBuilder struct {
	ca           *string
	clientID     *string
	clientSecret *string
	url          *string
}

// NewGitlabIdentityProvider creates a new builder of 'gitlab_identity_provider' objects.
func NewGitlabIdentityProvider() *GitlabIdentityProviderBuilder {
	return new(GitlabIdentityProviderBuilder)
}

// CA sets the value of the 'CA' attribute
// to the given value.
//
//
func (b *GitlabIdentityProviderBuilder) CA(value string) *GitlabIdentityProviderBuilder {
	b.ca = &value
	return b
}

// ClientID sets the value of the 'client_ID' attribute
// to the given value.
//
//
func (b *GitlabIdentityProviderBuilder) ClientID(value string) *GitlabIdentityProviderBuilder {
	b.clientID = &value
	return b
}

// ClientSecret sets the value of the 'client_secret' attribute
// to the given value.
//
//
func (b *GitlabIdentityProviderBuilder) ClientSecret(value string) *GitlabIdentityProviderBuilder {
	b.clientSecret = &value
	return b
}

// URL sets the value of the 'URL' attribute
// to the given value.
//
//
func (b *GitlabIdentityProviderBuilder) URL(value string) *GitlabIdentityProviderBuilder {
	b.url = &value
	return b
}

// Build creates a 'gitlab_identity_provider' object using the configuration stored in the builder.
func (b *GitlabIdentityProviderBuilder) Build() (object *GitlabIdentityProvider, err error) {
	object = new(GitlabIdentityProvider)
	if b.ca != nil {
		object.ca = b.ca
	}
	if b.clientID != nil {
		object.clientID = b.clientID
	}
	if b.clientSecret != nil {
		object.clientSecret = b.clientSecret
	}
	if b.url != nil {
		object.url = b.url
	}
	return
}
