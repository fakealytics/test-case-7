import("fmt")

package logger // TH

func log(message string) string {
	fmt.Println(message)

	return message;
}