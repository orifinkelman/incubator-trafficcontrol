/*

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

package fixtures

import "github.com/apache/incubator-trafficcontrol/traffic_ops/client"

// Parameters returns a default ParamResponse to be used for testing.
func Parameters() *client.ParamResponse {
	return &client.ParamResponse{
		Response: []client.Parameter{
			client.Parameter{
				Name:        "location",
				Value:       "/foo/trafficserver/",
				ConfigFile:  "parent.config",
				LastUpdated: "2012-09-17 15:41:22",
			},
		},
	}
}
