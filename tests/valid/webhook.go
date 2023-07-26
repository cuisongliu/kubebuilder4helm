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

package cronjob

// +kubebuilder4helm:webhook:webhookVersions=v1,verbs=create;update,path=/validate-testdata-kubebuilder-io-v1-cronjob,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=testdata.kubebuiler.io,resources=cronjobs,versions=v1,name=validation.cronjob.testdata.kubebuilder.io,sideEffects=None,timeoutSeconds=10,admissionReviewVersions=v1;v1beta1
// +kubebuilder4helm:webhook:verbs=create;update,path=/validate-testdata-kubebuilder-io-v1-cronjob,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=testdata.kubebuiler.io,resources=cronjobs,versions=v1,name=validation.cronjob.testdata.kubebuilder.io,sideEffects=NoneOnDryRun,timeoutSeconds=10,admissionReviewVersions=v1;v1beta1
// +kubebuilder4helm:webhook:webhookVersions=v1,verbs=create;update,path=/mutate-testdata-kubebuilder-io-v1-cronjob,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=testdata.kubebuiler.io,resources=cronjobs,versions=v1,name=default.cronjob.testdata.kubebuilder.io,sideEffects=None,timeoutSeconds=10,admissionReviewVersions=v1;v1beta1,reinvocationPolicy=IfNeeded