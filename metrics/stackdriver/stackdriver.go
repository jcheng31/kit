// Package stackdriver sends metrics to Stackdriver using Google's
// cloud.google.com/go/monitoring/apiv3 SDK.
package stackdriver

// Stackdriver receives metrics observations and forwards them to Stackdriver.
// Create a Stackdriver object, use it to create metrics, and pass those metrics
// as dependencies to the components that will use them.
//
// To regularly report metrics to Stackdriver, use the WriteLoop helper method.
type Stackdriver struct {
}
