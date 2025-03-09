// 渲染历史记录
function renderHistory() {
    const historyList = document.getElementById('historyList');
    historyList.innerHTML = history.map(item => `
        <div class="p-4 bg-white/50 rounded-lg hover:bg-white/70 transition-all cursor-pointer">
            <div class="flex justify-between items-start mb-2">
                <h3 class="font-medium text-gray-800 line-clamp-1">${item.url}</h3>
                <span class="text-sm text-gray-500">${item.date}</span>
            </div>
        </div>
    `).join('');
}

// 开始总结
async function startSummarizing(url,type) {
    try {
        // 显示加载状态
        document.getElementById('loadingIndicator').classList.remove('hidden');
        document.getElementById('resultCard').classList.add('hidden');
        document.getElementById('resultErrorCard').classList.add('hidden');

        // 调用后端 API
        const response = await fetch('/api/summarize', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url,type })
        });

        const data = await response.json();

        if (data.error) {
            throw new Error(data.error);
        }

        // 更新当前总结内容
        currentSummary = `${data.summary}`;

        // 更新预览内容
        document.getElementById('preview-content').innerHTML = marked.parse(currentSummary);
        document.getElementById('markdown-text').textContent = currentSummary;

        // 更新历史记录
        history.unshift({
            url,
            date: new Date().toISOString().split('T')[0]
        });
        renderHistory();

        // 显示结果
        document.getElementById('resultCard').classList.remove('hidden');
    } catch (error) {
        document.getElementById('err-info').textContent = error.message;
        document.getElementById('resultErrorCard').classList.remove('hidden');
        if (error.message.includes("获取文章内容失败")) {
            document.getElementById('err-info').innerHTML += "<br>建议尝试手动复制文章内容提交总结。"
        }
    } finally {
        document.getElementById('loadingIndicator').classList.add('hidden');
    }
}

// 切换标签
function switchTab(tabId) {
    // 获取当前点击的标签所在的区域
    const clickedTab = document.querySelector(`[data-tab="${tabId}"]`);
    const tabArea = clickedTab.closest('[role="tablist"], [role="inputtablist"]');
    const isInputTab = tabArea.getAttribute('role') === 'inputtablist';

    // 更新标签状态，只更新同一区域内的标签
    tabArea.querySelectorAll('[data-tab]').forEach(tab => {
        const isActive = tab.getAttribute('data-tab') === tabId;
        tab.classList.toggle('border-blue-500', isActive);
        tab.classList.toggle('text-blue-600', isActive);
        tab.classList.toggle('border-transparent', !isActive);
        tab.classList.toggle('text-gray-500', !isActive);
    });

    // 更新内容显示，只更新对应区域的内容
    const contentArea = isInputTab ? document.querySelector('.input-tab-content') : document.querySelector('.tab-content');
    contentArea.querySelectorAll('.tab-pane').forEach(pane => {
        pane.classList.toggle('hidden', pane.id !== `${tabId}-content`);
    });
}

// 复制到剪贴板
async function copyToClipboard(text) {
    try {
        await navigator.clipboard.writeText(text);
        alert('已复制到剪贴板');
    } catch (error) {
        alert('复制失败：' + error.message);
    }
}