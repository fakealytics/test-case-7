import("fmt")

package logger

func log(message string) string {
	fmt.Println(message)

	return message;
}