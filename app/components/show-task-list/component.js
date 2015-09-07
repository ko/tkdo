import Ember from 'ember';

export default Ember.Component.extend({
    actions: {
        deleteTask: function(task) {
            this.sendAction('deleteTask', task);
            this.set('task', {});
        }
    } 
});
