// +build prod

package main

func init() {
  features = append(features, "Hey this is PROD SPEAKIN'")
}