const orders = document.getElementById('orders');
const avgCheck = document.getElementById('avgCheck');
const uplift = document.getElementById('uplift');

const ordersValue = document.getElementById('ordersValue');
const avgCheckValue = document.getElementById('avgCheckValue');
const upliftValue = document.getElementById('upliftValue');

const extraRevenue = document.getElementById('extraRevenue');
const extraProfit = document.getElementById('extraProfit');

function formatNumber(value) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(value));
}

function calculate() {
  const ordersNum = Number(orders.value);
  const avgCheckNum = Number(avgCheck.value);
  const upliftPercent = Number(uplift.value) / 100;

  const revenueLift = ordersNum * avgCheckNum * upliftPercent;
  const profitLift = revenueLift * 0.35;

  ordersValue.textContent = formatNumber(ordersNum);
  avgCheckValue.textContent = `${avgCheckNum.toFixed(1)} €`;
  upliftValue.textContent = `${Math.round(upliftPercent * 100)}%`;

  extraRevenue.textContent = `${formatNumber(revenueLift)} €`;
  extraProfit.textContent = `${formatNumber(profitLift)} €`;
}

[orders, avgCheck, uplift].forEach((el) => {
  el.addEventListener('input', calculate);
});

calculate();
