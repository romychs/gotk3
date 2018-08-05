// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

package gtk

func init() {
	buildVersion = GTK_3_22
}
