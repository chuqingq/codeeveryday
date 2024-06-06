import torch
import torch.nn as nn
import torch.optim as optim
from torch.utils.data import DataLoader
from torchvision.datasets import MNIST
from torchvision.transforms import ToTensor

class SimpleNN(nn.Module):
    def __init__(self):
        super(SimpleNN, self).__init__()
        self.layer1 = nn.Linear(784, 256)
        self.layer2 = nn.Linear(256, 10)

    def forward(self, x):
        x = torch.relu(self.layer1(x))
        x = self.layer2(x)
        return x

# 创建神经网络实例
model = SimpleNN()

# 定义损失函数和优化器
loss_func = nn.CrossEntropyLoss()
optimizer = optim.SGD(model.parameters(), lr=0.01)

# 加载 MNIST 数据集
train_dataset = MNIST('./data', train=True, download=True, transform=ToTensor())
test_dataset = MNIST('./data', train=False, download=True, transform=ToTensor())

train_loader = DataLoader(train_dataset, batch_size=32, shuffle=True)
test_loader = DataLoader(test_dataset, batch_size=32, shuffle=False)

# 训练循环
model.train()
for epoch in range(10):
    for batch, (images, labels) in enumerate(train_loader):
        images = images.view(images.shape[0], -1)
        # print(f'{images.shape}')
        # 前向传播
        outputs = model(images)
        # 计算损失
        loss = loss_func(outputs, labels)
        # 反向传播
        optimizer.zero_grad()
        loss.backward()
        # 更新参数
        optimizer.step()

    # 在测试集上进行评估
    model.eval()
    with torch.no_grad():
        correct = 0
        total = 0
        for images, labels in test_loader:
            images = images.view(images.shape[0], -1)
            outputs = model(images)
            _, predicted = torch.max(outputs, 1)
            correct += (predicted == labels).sum().item()
            total += labels.size(0)
        accuracy = correct / total

    print(f'Epoch {epoch}: Loss {loss.item()}, Accuracy {accuracy}')

