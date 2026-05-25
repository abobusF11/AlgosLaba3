package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"algosLaba3/internal/usecase"
)

type Handler struct {
	validator *usecase.Validator
}

func NewHandler(validator *usecase.Validator) *Handler {
	return &Handler{
		validator: validator,
	}
}

func (h *Handler) Run(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	fmt.Fprintln(out, "Введите строки для проверки.")
	fmt.Fprintln(out, "Пустая строка завершает ввод.")
	fmt.Fprintln(out)

	for {
		fmt.Fprint(out, "> ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			break
		}

		result := h.validator.Validate(line)

		fmt.Fprintf(out, "%q -> %s\n", result.Input, result.Status())

		if !result.Valid {
			fmt.Fprintf(out, "Ошибка: позиция %d, %s\n", result.ErrorPos, result.ErrorReason)
		}

		fmt.Fprintf(out, "Пар: %d, глубина: %d, открывающих: %d, закрывающих: %d\n",
			result.PairsCount,
			result.MaxDepth,
			result.OpenCount,
			result.CloseCount,
		)

		fmt.Fprintln(out)
	}

	return scanner.Err()
}
