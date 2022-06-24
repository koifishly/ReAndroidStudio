.class public Landroid/koifishly/log/Logger;
.super Ljava/lang/Object;
.source "Logger.java"


# static fields
.field public static tag:Ljava/lang/String;


# direct methods
.method static constructor <clinit>()V
    .registers 1

    .prologue
    .line 6
    const-string v0, "ReverseStudio"

    sput-object v0, Landroid/koifishly/log/Logger;->tag:Ljava/lang/String;

    return-void
.end method

.method public constructor <init>()V
    .registers 1

    .prologue
    .line 5
    invoke-direct {p0}, Ljava/lang/Object;-><init>()V

    return-void
.end method

.method public static LogString(Ljava/lang/String;)V
    .registers 2
    .param p0, "str"    # Ljava/lang/String;

    .prologue
    .line 11
    const-string v0, "ReverseStudio"

    invoke-static {v0, p0}, Landroid/util/Log;->i(Ljava/lang/String;Ljava/lang/String;)I

    .line 12
    return-void
.end method

.method public static SetTag(Ljava/lang/String;)V
    .registers 1
    .param p0, "newTag"    # Ljava/lang/String;

    .prologue
    .line 8
    sput-object p0, Landroid/koifishly/log/Logger;->tag:Ljava/lang/String;

    .line 9
    return-void
.end method
