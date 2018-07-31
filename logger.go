import("fmt")

package logger // TH

func logger(message string) string {
	fmt.Println(message)

	return message;
}