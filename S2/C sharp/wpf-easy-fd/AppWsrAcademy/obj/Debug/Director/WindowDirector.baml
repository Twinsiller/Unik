   M S B A M L   `   `   `  ����H  DAppWsrAcademy, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null?$clr-namespace:AppWsrAcademy.DirectorAppWsrAcademy.Director  5   	   +    %AppWsrAcademy.Director.WindowDirector   5   	   R NWindowsBase, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35W SPresentationCore, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35\ XPresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35D 9http://schemas.microsoft.com/winfx/2006/xaml/presentation    5   	   R NSystem.Xaml, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089:x,http://schemas.microsoft.com/winfx/2006/xaml     5   	   @d2http://schemas.microsoft.com/expression/blend/2008     5   	   Jmc;http://schemas.openxmlformats.org/markup-compatibility/2006     5   	   6local$clr-namespace:AppWsrAcademy.Director     5   	   -   5          � WindowStyle$
  None=�5   	    � 
ResizeMode$ CanMinimize=�5
   	    � Title$! Окно директора��5   	   $	��450��6    $	��800��6-   .��5      �  � RowDefinitions 5   
   �� 5      $	��55* �6   �� 5      $
��364* �6   .+�5   
   �� $
��Left=�6   $	��Top=�5       �� Orientation$ 
Horizontal=�5      $��0q�5      .+�5      �� $ 
Horizontal=�6   $	��Top=�5      $��0q�5      .+�5      �� -   5        $��
BtnExitApp��5         BtnPanel#��[  5      ��5      �� 5      .R�5      ��  �� CornerRadius$ 20��5   !   
���  ��$��1q�5!   !    J� Opacity$	 0.7Q�5"   !   
����p��5#   !   5$      5&      �� 5'      -   5        $��BtnUpDownResize��5(      $��0,5,0,5q�5*      #��[  5+      ��5,      �� 5-      .R�5.      �� $ 20��5/   !   
��� � �
������$��1q�52   !   $	 0.7Q�53   !   55      57      �� 58      -   5        $��BtnMin��59      #��[  5;      ��5<      �� 5=      .R�5>      �� $ 20��5?   !   
��� ���$��1q�5A   !   $	 0.7Q�5B   !   
��� ���5C   !   5D      5F      5G      �� 5I      $��5q�5J      $��../Resources/wsr-300.png��5K      �� 5L      -   5        $��TxtName��5P      $
��12pt(�5M      $��DemiBold"�5N      $��Имя��5O      $��Center=�5Q      5R      �� 5S   
   $��1��6   .+�5V   
   &   System.Windows.Controls.DataGrid  -   5        $��GridListHistory��5W      $	��240��5V         AutoGenerateColumns .  "   VerticalScrollBarVisibility$
 Auto=�5Y      	   
IsReadOnly	 . 
   Columns
 5[      0  *System.Windows.Controls.DataGridTextColumn  5\      ,  &System.Windows.Controls.DataGridColumn   Width5  /System.Windows.Controls.DataGridLengthConverter$ * 5]         Header$  Преподаватель��5^      1  +System.Windows.Controls.DataGridBoundColumn   Binding 5_      �� *	User.Name+  5`      $ * 5a      $ Студент��5b       5c      �� *Student.Name+  5d      $ * 5e      $ Активность��5f       5g      �� *Status.Name+  5h      $ * 5i      $ Дата��5j       5k      �� *	DateEvent+ �� StringFormat$ 
dd.MM.yyyy��5m      �� 5o      $ 
Horizontal=�6   $��0,5,0,0q�5p      .+�5q      �� $ ��Записей итого:��6   $
��12pt(�5r      �� 5s      -   5        $��TxtTotal��5x      $
��12pt(�5t      $��Italic$�5u      $��5,0,0,0q�5v      $��количество��5w      5x      �� 5y      $
��12pt(�5z      $��Italic$�5{      $��5,0,0,0q�5|      $��стр.��5}      5~      �� 5      $ 
Horizontal=�6   $��0,5,0,0q�5�      .+�5�      �� $,��&Последнее изменение:��6   $
��12pt(�5�      �� 5�      -   5        $��TxtDate��5�      $
��12pt(�5�      $��Italic$�5�      $��5,0,0,0q�5�      $��дата��5�      5�      5�      �� 5�      $ 
Horizontal=�6   .+�5�      �� $��Фильтр:��6   $��DemiBold"�5�      $
��12pt(�5�      	� 5�      -	   5        $��RbDes��5�      $��По убыванию��5�      $
��12pt(�5�      $��Center=�5�      5�      	� 5�      -
   5        $
��RbUp��5�      $!��По возрастанию��5�      $
��12pt(�5�       d� 	IsChecked$
 TrueY�5�      $��Center=�5�      5�      5�      5�      5�      5�      