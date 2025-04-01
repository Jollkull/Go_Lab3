document.addEventListener('DOMContentLoaded', function() {
    const calculator1 = document.getElementById('calculator1');
    if (calculator1) {
        calculator1.addEventListener('submit', async function(e) {
            e.preventDefault();
            
            // Получаем все значения полей ввода
            const avgPower = parseFloat(document.getElementById('input1').value);
            const deviation = parseFloat(document.getElementById('input2').value);
            const improvedDeviation = parseFloat(document.getElementById('input3').value);
            const energyCost = parseFloat(document.getElementById('input4').value);
            
            try {
                const response = await fetch('/api/calculator1', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ 
                        values: [avgPower, deviation, improvedDeviation, energyCost]
                    }),
                });
                
                const data = await response.json();
                
                // Форматируем результат с переносами строк
                const formattedResult = data.result.replace(/\n/g, '<br>');
                document.getElementById('result').innerHTML = formattedResult;
                
            } catch (error) {
                console.error('Помилка:', error);
                document.getElementById('result').textContent = 'Сталася помилка під час розрахунку';
            }
        });
    }
});