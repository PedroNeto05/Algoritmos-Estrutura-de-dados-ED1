package algoritmos

import (
	"testing"
)

func TestIsValidParenteses(t *testing.T) {
	// Tabela com dezenas de cenários de teste
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Cenários de Sucesso (Validos)
		{"Vazio", "", true},
		{"Par simples", "()", true},
		{"Vários pares simples", "()()()", true},
		{"Aninhamento simples", "((()))", true},
		{"Misto e complexo", "(()(()))", true},
		{"Aninhamento profundo", "(((((((((())))))))))", true},
		{"Dois blocos aninhados", "((()))((()))", true},

		// Cenários de Falha (Invalido - Faltando pares)
		{"Apenas abertura", "(", false},
		{"Múltiplas aberturas", "(((", false},
		{"Esqueceu de fechar um", "(()", false},
		{"Faltando o último fechamento", "(((((((((()))))))))", false},

		// Cenários de Falha (Invalido - Fechamentos sobrando ou fora de ordem)
		{"Apenas fechamento", ")", false},
		{"Múltiplos fechamentos", ")))", false},
		{"Fechamento sobrando no fim", "())", false},
		{"Começando com fechamento", ")(", false},
		{"Invertido no meio", "())(()", false},
		{"Fechamento sobrando no meio", "())(", false},

		// Cenários de Falha (Invalido - Total igual, mas ordem errada)
		{"Total igual mas cruza", "))))((((", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Defer com recover para garantir que a suíte continue rodando
			// mesmo se a sua função IsValidParenteses der panic()
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Entrada '%s' causou panic: %v (Esperava retornar %v)", tt.input, r, tt.expected)
				}
			}()

			// Executa o seu algoritmo
			got := IsValidParenteses(tt.input)

			// Valida o resultado
			if got != tt.expected {
				t.Errorf("Entrada '%s' | Esperado: %v | Obtido: %v", tt.input, tt.expected, got)
			}
		})
	}
}
