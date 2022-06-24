#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    // 如果是Mac系统下，设置Mac风格的菜单栏
    ui->actionOpenFile->setMenuRole(QAction::PreferencesRole);
}

MainWindow::~MainWindow()
{
    delete ui;
}

/*
 * 打开一个apk，然后对这个apk进行解包。
 * 把结果显示到列表上。
 */
void MainWindow::on_actionOpenFile_triggered()
{
    // 1. 读取文件
}

