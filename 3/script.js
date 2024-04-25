fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1')
    .then(response => response.json())
    .then(data => {
        const table = document.getElementById('currency-table');
        data.forEach(currency => {
            const row = table.insertRow();
            row.insertCell(0).textContent = currency.id;
            row.insertCell(1).textContent = currency.symbol;
            row.insertCell(2).textContent = currency.name;
            if (currency.symbol === 'usdt') {
                row.style.background = 'green';
            } else if (data.indexOf(currency) < 5) {
                row.style.background = 'blue';
            }
        });
    })
    .catch(error => {
        console.error('Ошибка при получении данных:', error);
    });

